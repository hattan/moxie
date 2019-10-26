package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
  "simple-tf-provider/client"
)





func resourceTodoList() *schema.Resource {
  return &schema.Resource{
    Create: resourceTodoListCreate,
    Read:   resourceTodoListRead,
    Update: resourceTodoListUpdate,
    Delete: resourceTodoListDelete,

    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "path": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "category": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
        ValidateFunc: validation.StringInSlice([]string{
          string(todo_client.ShoppingList),
          string(todo_client.Home),
          string(todo_client.Office),
        }, false),
      },
    },
  }
}

func resourceTodoListCreate(d *schema.ResourceData, m interface{}) error {
  base_path := d.Get("path").(string)
  name := d.Get("name").(string)
  category := d.Get("category").(string)
  category_path := base_path + "/" + category
  file_path := category_path + "/" + name + ".txt"
  d.SetId(file_path)

  err := todo_client.Createfolder(base_path)
  if err != nil {
    return err
  }

  err = todo_client.Createfolder(category_path + "")
  if err != nil {
    return err
  }

  err = todo_client.CreateFile(file_path)
  if err != nil {
    return err
  }

  return resourceTodoListRead(d, m)
}

func resourceTodoListRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceTodoListUpdate(d *schema.ResourceData, m interface{}) error {
  // Enable partial state mode
  d.Partial(true)

  if d.HasChange("name") {
    // Try updating the address
    if err := todo_client.UpdateAddress(d, m); err != nil {
      return err
    }

    d.SetPartial("name")
  }

  // If we were to return here, before disabling partial mode below,
  // then only the "address" field would be saved.

  // We succeeded, disable partial mode. This causes Terraform to save
  // all fields again.
  d.Partial(false)

  return resourceTodoListRead(d, m)
}

func resourceTodoListDelete(d *schema.ResourceData, m interface{}) error {
  err := todo_client.DeleteFile(d.Id())
  if err != nil {
    return err
  }
  return nil
}
