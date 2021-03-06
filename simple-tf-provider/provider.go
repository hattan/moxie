package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    ResourcesMap: map[string]*schema.Resource{
      "todo_list": resourceTodoList(),
      "todo_storage" : resourceTodoStorage(),
      "todo_list_item" : resourceTodoListItem(),
    },
  }
}
