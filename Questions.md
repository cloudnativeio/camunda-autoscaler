- How long did it take you to solve the exercise? (Please be honest, we evaluate your answer to this question based on your experience.)
  - It took me 6 hours to complete the exercise

- Which additional steps would you take in order to make this code production ready? Why?
  - Adding unit test on the code and hardening the error management. The code doesn't have a unit test and go routine management has not been taken into consideration.

- Which step took most of the time? Why?
  - Figuring out why minikube doesn't run on my workstation anymore. I found out that due to the security hardening, it doesn't allow me to access the proxy. I have to instantiate a VM in gcp to and install needed binaries like kubectl and minikube.