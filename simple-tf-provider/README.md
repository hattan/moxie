# Terraform ToDoList Provider

## This is not a real project, but an exercise in order to learn how to build [Terraform Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html)

The goal is to create a ToDo list that is stored on disk using files and folders.

```ruby 
resource "todo_list" "project_one" {
   name = "TestProject"
   base_path= "~/todos"
   category = "Home"
}

resource "todo_list_item" "project_one" {
   description = "Go for a walk"
   list_name = todo_list.name
}
```
### Running the Project:
* Ensure you have [Terraform](https://www.terraform.io) and [GoLang](https://golang.org/) installed.
* Invoke [build.sh](./build.sh)
* Navigate to the [infra folder](./infra).
* ```terraform apply -auto-approve ```
* Verify that a file called TestProject.txt was created in base_path/category.
* ```terraform delete -auto-approve ```
* Verify that file was deleted.

### Project Status:
* todo_list resource is now supported. A file with the project name is created under the base_path / category.

### ToDo:
* Implement todo_list_item.
* Add Unit Tests.
* Refactor base folder into it's own resource so it can be managed via terraform as well.
* Add support Updates..
* Create step by step walkthrough.


Resources:
* [Terraform.io](https://www.terraform.io)
* [Writing Custom Terraform Providers](https://www.terraform.io/docs/extend/writing-custom-providers.html)
* [Terraform Schema Types](https://www.terraform.io/docs/extend/schemas/schema-types.html)
* [ARM Bot Web App Resource](https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/azurerm/resource_arm_bot_web_app.go)
* [Debugging Terraform](https://www.terraform.io/docs/internals/debugging.html)
* [mage - a rake like build tool for go](https://github.com/magefile/mage)
