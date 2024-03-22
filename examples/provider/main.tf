terraform {
  required_providers {
    liblab = {
      source  = "hashicorp.com/edu/liblab"
      version = "0.0.1"
    }
  }
}

provider "liblab" {

  host = "https://api.example.com"

  auth_token = "MY_TOKEN"

}
