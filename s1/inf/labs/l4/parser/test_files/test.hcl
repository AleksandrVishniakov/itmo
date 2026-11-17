// --- Comment 1 ---
/* --- Comment 2 */
# --- Comment 3 ---

/*
multiline
***
comment
*/


# 1. Блок без меток, содержит только атрибуты (контейнер)
settings-1 {
  enabled     = true
  retries     = 3e-8
  ratio       = -0.75
  name        = "example"
  description = "Простой набор атрибутов \"ЦИТАТА\""
}

# 2. Блок с ОДНОЙ меткой
service "frontend" {
  port    = "80:80"
  protocol = null

  # 3. Вложенный блок без меток (тело — объект)
  healthcheck {
    path     = "/health"
    interval = 30
  }

  # 4. Повторяемый вложенный блок (должен стать массивом)
  tag {
    key   = "env"
    value = (a-b) >= (1+2-3 * 4 / 5) *5/6.1 || port.len != "10"
  }
  tag {
    key   = "team"
    value = "web"
  }
}

# 5. Блок с ДВУМЯ метками (type, name)
resource "compute" "web" {
  # 6. Выражение / интерполяция
  image_id  = var.image_id
  region    = "${var.region}-1"
  user_data = <<-EOF
    #!/bin/bash
    echo "hello"
  EOF

  # 7. Список (list) примитивов
  ports = [80, 443, 8080]

  # 8. Объект (map) как значение атрибута
  labels = {
    app     = "web", x="y"
    version = "1.2.3"
  }

  # 9. Список объектов
  rules = [
    { port = 80, proto = "tcp" },
    { port = 443, proto = "tcp" }
  ]

  # 10. Повторяемый блок с меткой (массив вложенных объектов)
  disk "data" {
    size = 100
  }
  disk "logs" {
    size = 50
  }

  # 11. Вложенный блок, содержащий повторяемый блок
  network {
    ingress {
      from = 22
      to   = 22
    }
    ingress {
      from = 80
      to   = 80
    }
  }
}

# 12. Блок с меткой, где метка — часть "типа" на верхнем уровне
variable "region" {
  type    = "string"
  default = "eu-central-1"
}

variable "image_id" {
  type = "string"
}
 
