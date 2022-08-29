# Attributes

Define characteristics of audiences that decide the experiences they will be included in.

## Example Usage

```hcl
resource "optimizely_attribute" "country_attribute" {
    key = "COUNTRY"
    archived = false
    name = "country"
    project_id = 123456789
    description = "Attribute used to define a country"
}
```

## Argument Reference

* `archived` - Whether or not the Attribute has been archived.
* `description` - A short description of the Attribute.
* `key` - Unique string identifier for this Attribute within the project.
* `name` - The name of the Attribute. For Full Stack projects, the name will be set to the value of the key.
* `project_id` - The ID of the project the Attribute belongs to.  
