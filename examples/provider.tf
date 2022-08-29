terraform {
  required_providers {
    optimizely = {
      source  = "pffreitas/optimizely"
      version = "0.10"
    }
  }
}

provider "optimizely" {
  host  = "https://api.optimizely.com"
  token = var.api_token
}
