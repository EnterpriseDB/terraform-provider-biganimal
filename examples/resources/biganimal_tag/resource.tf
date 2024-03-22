terraform {
  required_providers {
    biganimal = {
      source  = "EnterpriseDB/biganimal"
      version = "0.8.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.0"
    }
  }
}

resource "biganimal_tag" "tag" {
  tag_name = "tag-name"
  // The default colors to choose from are:
  // magenta, red, orange, yellow, green, teal, blue, indigo, purple and grey
  // you can also use a custom hex color code e.g. #e34545
  color   = "blue"
}