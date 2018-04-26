## of_badge_gen - status badge for opendfaas-cloud

[![OpenFaaS](https://img.shields.io/badge/openfaas-serverless-blue.svg)](https://www.openfaas.com)
[![Status](http://0341c281.ngrok.io/function/of_badge_gen?user=s8sg&repo=of_badge_gen&branch=master)](http://0341c281.ngrok.io/of_badge_gen)



#### About
of_badge_gen is a faas function to generate status badge for function deployed in opendfaas-cloud 

[![success_DEPLOY](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/success.svg)](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/success.svg)   -  Successfully Deployed  

[![pending_BUILD](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/pending_BUILD.svg)](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/pending_BUILD.svg)   -  Pending at Build   

[![failure_BUILD](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/failure_BUILD.svg)](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/failure_BUILD.svg)   -  Failed to Build  

[![pending_DEPLOY](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/pending_DEPLOY.svg)](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/pending_DEPLOY.svg)   -  Pending at Deploy  

[![failure_DEPLOY](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/failure_DEPLOY.svg)](https://raw.githubusercontent.com/s8sg/of_badge_gen/master/assets/image/failure_DEPLOY.svg)   -  Failed to Deploy  


#### Deploying in your openfaas-cloud

1. Fork this repo
2. Install your openfaas-cloud github app on the forked repo
3. Change the CUSTOMER url, it should be same that is used in openfaas-cloud
4. Use the badge for a repo as:
   ```
   [![Status](http://your.openfaas.cloud/function/of_badge_gen?user=<repo_user_or_org>&repo=<repo_name>&branch=master)](repo_url)
   ```
