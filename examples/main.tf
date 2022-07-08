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
  secret_key = "h3110_w0r1d"
}

resource "filesystem_file" "writer" {
  path    = "./helloworld.txt"
  content = "你好世界, Hello World"
}

data "filesystem_file" "reader" {
  path = filesystem_file.writer.path
}

output "file_details" {
  value = <<EOH
  name:    ${data.filesystem_file.reader.filename}
  size:    ${data.filesystem_file.reader.size}
  update:  ${data.filesystem_file.reader.update_time}
  content: ${data.filesystem_file.reader.content}
EOH
}
