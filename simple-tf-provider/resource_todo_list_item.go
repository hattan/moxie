package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


func resourceTodoListItem() *schema.Resource {
  return &schema.Resource{
    Create: resourceTodoListItemCreate,
    Read:   resourceTodoListItemRead,
    Update: resourceTodoListItemUpdate,
    Delete: resourceTodoListItemDelete,

    Schema: map[string]*schema.Schema{
      "description": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "list_name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "list_category": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },      
      "path": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },            
    },
  }
}

func resourceTodoListItemCreate(d *schema.ResourceData, m interface{}) error {
  description := d.Get("description").(string)
  base_path := d.Get("path").(string)
  list_name := d.Get("list_name").(string)
  list_category := d.Get("list_category").(string)
  category_path := base_path + "/" + list_category
  file_path := category_path + "/" + list_name + ".txt"
  

  d.SetId(file_path + "[" + description + "]")

  return resourceTodoListItemRead(d, m)
}

func resourceTodoListItemRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceTodoListItemUpdate(d *schema.ResourceData, m interface{}) error {
  return resourceTodoListItemRead(d, m)
}

func resourceTodoListItemDelete(d *schema.ResourceData, m interface{}) error {
  return nil
}
