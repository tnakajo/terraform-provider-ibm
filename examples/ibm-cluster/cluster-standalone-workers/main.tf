resource "ibm_container_cluster" "testacc_cluster" {
  name            = "mycluster0705"
  datacenter      = "dal10"
  #machine_type    = "b3c.4x16"
  hardware        = "shared"
  #public_vlan_id  = "123456"
  #private_vlan_id = "654321"
  #subnet_id       = ["1154643"]
  #default_pool_size      = 1
}
