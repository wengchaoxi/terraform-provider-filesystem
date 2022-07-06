terraform {
  required_providers {
    filesystem = {
      version = "0.0.1"
      source  = "github.com/wengchaoxi/filesystem"
    }
  }
}

provider "filesystem" {
  # Or use FILESYSTEM_SECRET_KEY env var
  secret_key = "secret_key_helloworld"
}

resource "filesystem_file" "this" {
  path    = "./helloworld.txt"
  content = "你好世界, Hello World"
}

data "filesystem_file" "this" {
  path = filesystem_file.this.id
}

output "file_details" {
  value = <<EOH
  name:    ${data.filesystem_file.this.filename}
  size:    ${data.filesystem_file.this.size}
  update:  ${data.filesystem_file.this.update_time}
  content: ${data.filesystem_file.this.content}
EOH
}
