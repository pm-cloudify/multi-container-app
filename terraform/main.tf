terraform {
  required_providers {
    arvan = {
      source = "terraform.arvancloud.ir/arvancloud/iaas"
    }
  }
}

provider "arvan" {
  api_key = var.api_key
}

data "arvan_images" "terraform_image" {
  region     = var.region
  image_type = "distributions"
}


data "arvan_plans" "plan_list" {
  region = var.region
}

data "arvan_server_groups" "server_group_list" {
  region = var.region
}

data "arvan_dedicated_servers" "terraform_dedicated_server" {
  region = var.region
}

locals {
  chosen_image = try(
    [for image in data.arvan_images.terraform_image.distributions : image
    if image.distro_name == var.chosen_distro_name && image.name == var.chosen_name],
    []
  )

  selected_plan = [for plan in data.arvan_plans.plan_list.plans : plan if plan.id == var.chosen_plan_id][0]
}

resource "arvan_security_group" "terraform_security_group" {
  region      = var.region
  description = "Terraform-created security group"
  name        = "tf_security_group"
  rules = [
    {
      direction = "ingress"
      protocol  = "icmp"
    },
    {
      direction = "ingress"
      protocol  = "udp"
    },
    {
      direction = "ingress"
      protocol  = "tcp"
    },
    {
      direction = "egress"
      protocol  = ""
    }
  ]
}

resource "arvan_volume" "terraform_volume" {
  region      = var.region
  description = "Terraform-created volume"
  name        = "tf_volume"
  size        = 9
}

data "arvan_networks" "terraform_network" {
  region = var.region
}

locals {
  chosen_dedicated_server = try(
    [for ds in data.arvan_dedicated_servers.terraform_dedicated_server.dedicated_servers : ds if contains(ds.labels, var.dedicated_server_label)], []
  )
}

resource "arvan_network" "terraform_private_network" {
  region      = var.region
  description = "Terraform-created private network"
  name        = "tf_private_network"
  dhcp_range = {
    start = "10.255.255.19"
    end   = "10.255.255.150"
  }
  dns_servers    = ["8.8.8.8", "1.1.1.1"]
  enable_dhcp    = true
  enable_gateway = true
  cidr           = "10.255.255.0/24"
  gateway_ip     = "10.255.255.1"
}

resource "arvan_abrak" "built_by_terraform" {
  depends_on = [arvan_volume.terraform_volume, arvan_network.terraform_private_network, arvan_security_group.terraform_security_group]
  timeouts {
    create = "1h30m"
    update = "2h"
    delete = "20m"
    read   = "10m"
  }
  region          = var.region
  name            = "built_by_terraform"
  count           = 1
  image_id        = local.chosen_image[0].id
  flavor_id       = local.selected_plan.id
  disk_size       = 25
  server_group_id = var.chosen_server_group_id
  enable_ipv4     = true
  enable_ipv6     = true
  networks = [
    {
      network_id = arvan_network.terraform_private_network.network_id
    }
  ]
  security_groups = [arvan_security_group.terraform_security_group.id]
  volumes         = [arvan_volume.terraform_volume.id]

  # add your define ssh key name from Arvan panel
  # ssh_key_name = 
}

output "instances" {
  value = arvan_abrak.built_by_terraform
}
