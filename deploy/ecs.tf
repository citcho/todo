resource "aws_ecr_repository" "app" {
  provider             = aws.todo
  name                 = "app"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecs_cluster" "app" {
  provider = aws.todo
  name     = "${local.project}-${local.env}-cluster"
}

resource "aws_ecs_service" "app" {
  provider        = aws.todo
  name            = "${local.project}-${local.env}-service"
  cluster         = aws_ecs_cluster.app.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = module.vpc.public_subnets
    security_groups  = [aws_security_group.app_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.alb_target_group.arn
    container_name   = "app"
    container_port   = tonumber(data.aws_ssm_parameter.app_port.value)
  }
}

resource "aws_ecs_task_definition" "app" {
  provider                 = aws.todo
  family                   = "app"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 1024
  memory                   = 2048
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.app.arn
  container_definitions = jsonencode([
    {
      name      = "app"
      image     = "${aws_ecr_repository.app.repository_url}:latest"
      essential = true
      cpu       = 512
      memory    = 1024
      portMappings = [
        {
          containerPort = tonumber(data.aws_ssm_parameter.app_port.value)
          hostPort      = 80
        }
      ]
      environment = [
        {
          name  = "CLIENT_HOST"
          value = tostring(data.aws_ssm_parameter.app_client_host.value)
        },
        {
          name  = "CLIENT_PORT"
          value = tostring(data.aws_ssm_parameter.app_client_port.value)
        },
        {
          name  = "TODO_PORT"
          value = tostring(data.aws_ssm_parameter.app_port.value)
        },
        {
          name  = "DB_NAME"
          value = tostring(data.aws_ssm_parameter.db_name.value)
        },
        {
          name  = "DB_USER"
          value = tostring(data.aws_ssm_parameter.db_username.value)
        },
        {
          name  = "DB_PASS"
          value = tostring(data.aws_ssm_parameter.db_password.value)
        },
        {
          name  = "DB_HOST"
          value = tostring(aws_db_instance.app_db.address)
        },
        {
          name  = "DB_PORT"
          value = tostring(aws_db_instance.app_db.port)
        },
        {
          name  = "TZ"
          value = tostring(data.aws_ssm_parameter.app_timezone.value)
        },
        {
          name  = "BUNDEBUG"
          value = tostring(data.aws_ssm_parameter.app_bundebug.value)
        },
      ]
      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-create-group  = "true",
          awslogs-group         = tostring(aws_cloudwatch_log_group.app.name),
          awslogs-region        = local.region,
          awslogs-stream-prefix = "ecs"
        }
      }
    },
  ])
}

resource "aws_iam_role" "app" {
  provider = aws.todo
  name     = "${local.project}-${local.env}-app-task-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com",
        },
        Action = "sts:AssumeRole",
      },
    ],
  })

  inline_policy {
    name = "${local.project}-${local.env}-app-task-policy"
    policy = jsonencode(
      {
        Statement = [
          {
            Action = [
              "ssm:GetParameters",
            ]
            Effect = "Allow"
            Resource = [
              data.aws_ssm_parameter.app_client_host.arn,
              data.aws_ssm_parameter.app_client_port.arn,
              data.aws_ssm_parameter.app_port.arn,
              data.aws_ssm_parameter.db_name.arn,
              data.aws_ssm_parameter.db_username.arn,
              data.aws_ssm_parameter.db_password.arn,
              data.aws_ssm_parameter.db_port.arn,
              data.aws_ssm_parameter.app_timezone.arn,
              data.aws_ssm_parameter.app_bundebug.arn,
            ]
          },
        ]
        Version = "2012-10-17"
      }
    )
  }
  managed_policy_arns = [
    data.aws_iam_policy.ecs_task_execution_role_policy.arn,
    data.aws_iam_policy.ec2_container_registry_read_only.arn,
  ]
}

data "aws_iam_policy" "ecs_task_execution_role_policy" {
  provider = aws.todo
  name     = "AmazonECSTaskExecutionRolePolicy"
}

data "aws_iam_policy" "ec2_container_registry_read_only" {
  provider = aws.todo
  name     = "AmazonEC2ContainerRegistryReadOnly"
}
