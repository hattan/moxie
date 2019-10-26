resource "todo_storage" "local" {
   path= var.base_path
}

resource "todo_list" "project_one" {
   name = var.project_name
   path= todo_storage.local.path
   category = "Home"
}