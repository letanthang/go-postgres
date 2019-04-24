export APP_NAME=go-postgres
export DATE=`date +%s`
export RELEASE=staging
  
  function build() {
    # registry_login
    docker login -u letanthang -p "$HUB_DOCKER_PASS"
        echo "Building Dockerfile-based application..."
        #build image
        docker build \
          --build-arg HTTP_PROXY="$HTTP_PROXY" \
          --build-arg http_proxy="$http_proxy" \
          --build-arg HTTPS_PROXY="$HTTPS_PROXY" \
          --build-arg https_proxy="$https_proxy" \
          --build-arg FTP_PROXY="$FTP_PROXY" \
          --build-arg ftp_proxy="$ftp_proxy" \
          --build-arg NO_PROXY="$NO_PROXY" \
          --build-arg no_proxy="$no_proxy" \
          -t "letanthang/go-postgres" .
        
        docker push letanthang/go-postgres
      
  }
  function deploy() {
    #   sed 's/_VERSION_/'"latest"'/g; s/APP_NAME/'"$CI_PROJECT_NAME"'/g; s/_DATE_/'"$DATE"'/g; s/env-/'"$RELEASE-"'/g' ./provision/k8s/* > deployment.yaml
      sed 's/APP_NAME/'"$APP_NAME"'/g; s/_DATE_/'"$DATE"'/g; s/env-/'"$RELEASE-"'/g' ./provision/k8s/deployment.yaml > zdeployment.yaml
      kubectl apply -f zdeployment.yaml
  }
  
  build
  deploy


  # function install_dependencies() {
  #   #install curl
  #   apk add --no-cache curl
    
  #   # install kubectl
  #   curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
  #   chmod +x ./kubectl
  #   mv ./kubectl /usr/local/bin/kubectl
    
  #   #pre-config kubectl
  #   mkdir ~/.kube
  #   echo "$KUBE_STAGING" >> /tmp/kube_config
  #   base64 -d /tmp/kube_config >> ~/.kube/config
    
  #   #install helm
  #   curl "https://kubernetes-helm.storage.googleapis.com/helm-v${HELM_VERSION}-linux-amd64.tar.gz" | tar zx
  #   mv linux-amd64/helm /usr/bin/
  #   mv linux-amd64/tiller /usr/bin/
  #   helm version --client
  #   tiller -version
  # }
