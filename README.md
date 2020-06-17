# env-gae-run-template


    # Generate common dotenv file.
    APP_NAME=awesome GCP_PROJECT_DEV=my-dev-project RUN_REGION_DEV=us-central1 FOO_HOGE_DEV="dev-hoge" BAR_PIYO_DEV="dev-piyo" \
        bash gen-common-env.sh development
    APP_NAME=awesome GCP_PROJECT_STG=my-stg-project RUN_REGION_STG=us-east1 FOO_HOGE_STG="stg-hoge" BAR_PIYO_STG="stg-piyo" \
        bash gen-common-env.sh staging
    APP_NAME=awesome GCP_PROJECT_PRD=my-prd-project RUN_REGION_PRD=asia-northeast1 FOO_HOGE_PRD="prd-hoge" BAR_PIYO_PRD="prd-piyo"\
        bash gen-common-env.sh production

    # Generate application dotenv and yaml files.
    make gen-app-env
    make gen-app-yaml
    GAE_VERSION=test make deploy-gae
    make deploy-run


## env variables list

| env name | env variables                                                                  |
|----------|--------------------------------------------------------------------------------|
| dev      | APP_ENV, APP_NAME, FOO_HOGE_DEV, BAR_PIYO_DEV, GCP_PROJECT_DEV, RUN_REGION_DEV |
| stg      | APP_ENV, APP_NAME, FOO_HOGE_STG, BAR_PIYO_STG, GCP_PROJECT_STG, RUN_REGION_STG |
| prd      | APP_ENV, APP_NAME, FOO_HOGE_PRD, BAR_PIYO_PRD, GCP_PROJECT_PRD, RUN_REGION_PRD |

| app name | env variables               |
|----------|-----------------------------|
| api      | APP_ENV, APP_NAME           |
| foo      | APP_ENV, APP_NAME, APP_HOGE |
| bar      | APP_ENV, APP_NAME, APP_PIYO |
| baz      | APP_ENV, APP_NAME           |

## refs.

* [Cloud Run Region](https://cloud.google.com/run/docs/locations?hl=en)