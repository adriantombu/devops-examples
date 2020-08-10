# AWS S3 Static Hosting

A simple html page hosted on AWS S3. The infrastructure needed is created with [Terraform](https://www.terraform.io/), and the site is deployed with the [AWS Command Line Interface](https://aws.amazon.com/cli/).

### Steps to deploy

- `terraform init` : initialize Terraform  (it will download the AWS provider)
- `terraform plan` : check what resources will be created (optional)
- `terraform apply` : create the infrastructure on AWS
- `aws s3 sync . s3://s3-static-site.devops.otso.fr --exclude '*' --include '*.html'` : deploy the website to AWS S3
- Go to http://**YOUR-BUCKET-NAME**.s3-website-**YOUR-REGION**.amazonaws.com to see your website. In this example it would be `http://s3-static-site.devops.otso.fr.s3-website-eu-west-1.amazonaws.com/`

You can destroy all the resources created with the command `terraform destroy`
