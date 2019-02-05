{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Deploy application into Kubernetes.

Command will create Helm Release and wait until all resources of the release are become ready.

Deploy needs the same parameters as push to construct image names: repo and tags. Docker images 
names are constructed from paramters as IMAGES_REPO/IMAGE_NAME:TAG. Deploy will fetch built image 
ids from Docker repo. So images should be published prior running deploy.

Helm chart directory .helm should exists and contain valid Helm chart.

Environment is a required param for the deploy by default, because it is needed to construct Helm 
Release name and Kubernetes Namespace. Either --env or WERF_DEPLOY_ENVIRONMENT should be specified 
for command.

Read more info about Helm chart structure, Helm Release name, Kubernetes Namespace and how to 
change it: https://flant.github.io/werf/reference/deploy/deploy_to_kubernetes.html

{{ header }} Syntax

```bash
werf deploy [options]
```

{{ header }} Environments

```bash
  $WERF_SECRET_KEY     Use specified secret key to extract secrets for the deploy; recommended way 
                       to set secret key in CI-system
  $WERF_DOCKER_CONFIG  Force usage of the specified docker config
```

{{ header }} Options

```bash
      --dir='':
            Change to the specified directory to find werf.yaml config
      --env='':
            Use specified environment (use WERF_DEPLOY_ENVIRONMENT by default)
  -h, --help=false:
            help for deploy
      --home-dir='':
            Use specified dir to store werf cache files and dirs (use WERF_HOME environment or 
            ~/.werf by default)
  -i, --images='':
            Docker Repo to store images (use WERF_IMAGES_REPO environment by default)
  -p, --images-password='':
            Images Docker repo username (granted permission to read images info, use 
            WERF_IMAGES_PASSWORD environment by default)
  -u, --images-username='':
            Images Docker repo username (granted permission to read images info, use 
            WERF_IMAGES_USERNAME environment by default)
      --kube-context='':
            Kubernetes config context
      --namespace='':
            Use specified Kubernetes namespace (use %project-%environment template by default)
      --release='':
            Use specified Helm release name (use %project-%environment template by default)
      --secret-values=[]:
            Additional helm secret values
      --set=[]:
            Additional helm sets
      --set-string=[]:
            Additional helm STRING sets
      --ssh-key=[]:
            Enable only specified ssh keys (use system ssh-agent by default)
  -s, --stages='':
            Docker Repo to store stages or :local for non-distributed build (only :local is 
            supported for now)
      --stages-password='':
            Stages Docker repo password
      --stages-username='':
            Stages Docker repo username
      --tag=[]:
            Add tag (can be used one or more times)
      --tag-git-branch='':
            Tag by git branch (use WERF_AUTOTAG_GIT_BRANCH environment by default)
      --tag-git-commit='':
            Tag by git commit (use WERF_AUTOTAG_GIT_COMMIT environment by default)
      --tag-git-tag='':
            Tag by git tag (use WERF_AUTOTAG_GIT_TAG environment by default)
  -t, --timeout=0:
            Resources tracking timeout in seconds
      --tmp-dir='':
            Use specified dir to store tmp files and dirs (use WERF_TMP environment or system tmp 
            dir by default)
      --values=[]:
            Additional helm values
```
