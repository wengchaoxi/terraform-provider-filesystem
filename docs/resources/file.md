# filesystem_file (Resource)

## Example Usage

```hcl
resource "filesystem_file" "writer" {
  path    = "./helloworld.txt"
  content = "Hello World"
}
```

## Argument Reference

The following arguments are supported:

* `path` - (Required) The path to save the file.
* `content` - (Optional) The content of the file.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of this resource.
