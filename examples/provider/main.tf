terraform {
  required_providers {
    liblab = {
      source  = "hashicorp.com/kapicic/liblab"
      version = "0.0.7"
    }
  }
}

provider "liblab" {

  host = "https://api.example.com"

  auth_token = "MY_TOKEN"

}
