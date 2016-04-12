# s3gateway

A simple service that pre-signs S3 URLs.

When combined with access controls, this service can provide access to private
resources in S3 to clients without those clients being S3-aware. For a simple
example, suppose that we had a yum repository stored in a private S3 bucket.
Instead of making all our servers S3-aware, we could point all of our servers
to this gateway and use a firewall to ensure that only approved servers can
access the objects in our yum repo.

This gateway is intended to be deployed with a proxy server or firewall that
controls access to this service.

## Usage

Set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables or
put them into `~/.aws/credentials`. Set `AWS_REGION`.

Run `s3gateway BUCKET` to start the service on port 8080.

Request <localhost:8080/some-key-in-bucket> and the service will redirect you
to a pre-signed URL for the requested key in the specified bucket.
