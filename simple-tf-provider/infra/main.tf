resource "todo_storage" "local" {
   path= var.base_path
}

resource "todo_list" "project_one" {
   name = var.project_name
   path= todo_storage.local.path
   category = "Home"
}

resource "todo_list_item" "walk_todo" {
   description = "Go for a walk"
   list_name = todo_list.project_one.name
   list_category = todo_list.project_one.category
   path = todo_list.project_one.path
}