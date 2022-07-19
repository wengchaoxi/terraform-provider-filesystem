# filesystem_file (Data Source)

## Example Usage

```hcl
resource "filesystem_file" "writer" {
  path    = "./helloworld.txt"
  content = "Hello World"
}

data "filesystem_file" "reader" {
  path = filesystem_file.writer.path
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Required) The path to save the file.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `content` - The content of the file.
* `filename` - The name of the file.
* `id` - The ID of this resource.
* `size` - The size of The file.
* `update_time` - The update time of the file.
