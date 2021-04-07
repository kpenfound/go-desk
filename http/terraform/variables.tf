variable "region" {
  default = "us-east-1"
}
variable "app_version" {
  description = "Version of the function stored in s3 at {version}/function.zip"
  default     = "v0.0.0"
}
