resource "random_id" "test" {
  byte_length = 9
}

output "random_id" {
  value = random_id.test.id
}