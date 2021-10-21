package config

// Default bucket name for S3.
const DefaultS3Bucket = "gofi"

const defaultConfig = `
logger:
  level: debug
  path: "/var/log/gofi.log"
  max-size: 1024
  max-backups: 7
  max-age: 7
  stdout: true

server:
  address: :7677
  read-timeout: 20s
  write-timeout: 20s
  graceful-timeout: 5s

s3:
  endpoint: "127.0.0.1:9000"
  force-path-style: true
  access-key: "access-key"
  secret-key: "secret-key"
  session-token: ""
  disable-ssl: true
  region: "eu-east-1"
`
