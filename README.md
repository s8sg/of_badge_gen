## of_badge_gen
faas function to generate status badge for function deployed in opendfaas-cloud


#### Integration with your openfaas-cloud

1. Fork this repo
2. Install your openfaas-cloud github app on the forked repo
3. Change the CUSTOMER url, it should be same accross your openfaas-cloud
4. Use the badge for a repo as:   
   `http://your.openfaas.cloud/function/of_badge_gen?user=<repo_user_or_org>&repo=<repo_name>&branch=master`
