package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "simple-tf-provider/client"
)


func resourceTodoStorage() *schema.Resource {
  return &schema.Resource{
    Create: resourceTodoStorageCreate,
    Read:   resourceTodoStorageRead,
    Update: resourceTodoStorageUpdate,
    Delete: resourceTodoStorageDelete,

    Schema: map[string]*schema.Schema{
      "path": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
  }
}

func resourceTodoStorageCreate(d *schema.ResourceData, m interface{}) error {
  path := d.Get("path").(string)
	d.SetId(path)
	
  err := todo_client.Createfolder(path)
  if err != nil {
    return err
	}
	
  return resourceTodoStorageRead(d, m)
}

func resourceTodoStorageRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceTodoStorageUpdate(d *schema.ResourceData, m interface{}) error {
 	return resourceTodoStorageRead(d, m)
}

func resourceTodoStorageDelete(d *schema.ResourceData, m interface{}) error {
	path := d.Id()
	err := todo_client.RemoveContents(path)
	if err != nil {
    return err
	}

  err = todo_client.DeleteFile(path)
  if err != nil {
    return err
  }
  return nil
}
