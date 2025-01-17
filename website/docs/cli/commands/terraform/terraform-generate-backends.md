---
title: atmos terraform generate backends
sidebar_label: generate backends
sidebar_class_name: command
id: generate-backends
description: Use this command to generate the Terraform "backend config" files for all `atmos` Terraform Components in all Stacks.
---

:::note purpose
Use this command to generate the Terraform "backend config" files for all `atmos` Terraform [Components](/core-concepts/components) in all [Stacks](/core-concepts/stacks).
:::

## Usage

Execute the `terraform generate backends` command like this:

```shell
atmos terraform generate backends [options]
```

This command generates backend config files for all `atmos` terraform components in all stacks.

:::tip
Run `atmos terraform generate backends --help` to see all the available options
:::

## Examples

```shell
atmos terraform generate backends --file-template {component-path}/{tenant}/{environment}-{stage}.tf.json --format json
atmos terraform generate backends --file-template {component-path}/backends/{tenant}-{environment}-{stage}.tf.json --format json
atmos terraform generate backends --file-template backends/{tenant}/{environment}/{region}/{component}.tf --format hcl
atmos terraform generate backends --file-template backends/{tenant}-{environment}-{stage}-{component}.tf
atmos terraform generate backends --file-template /{tenant}/{stage}/{region}/{component}.tf
atmos terraform generate backends --file-template backends/{tenant}-{environment}-{stage}-{component}.tfbackend --format backend-config
atmos terraform generate backends --stacks orgs/cp/tenant1/staging/us-east-2,orgs/cp/tenant2/dev/us-east-2 --file-template <file_template>
atmos terraform generate backends --stacks tenant1-ue2-staging,tenant1-ue2-prod --file-template <file_template>
atmos terraform generate backends --stacks orgs/cp/tenant1/staging/us-east-2,tenant1-ue2-prod --file-template <file_template>
atmos terraform generate backends --components <component1>,<component2> --file-template <file_template>
atmos terraform generate backends --format hcl --file-template <file_template>
atmos terraform generate backends --format json --file-template <file_template>
atmos terraform generate backends --format backend-config --file-template <file_template>
```

## Flags

| Flag              | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Alias | Required |
|:------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:------|:---------|
| `--file-template` | Backend file template (path, file name, and file extension).<br/>Supports absolute and relative paths.<br/>Supports context tokens: `{namespace}`, `{tenant}`, `{environment}`,<br/>`{region}`, `{stage}`, `{base-component}`, `{component}`, `{component-path}`.<br/>All subdirectories in the path will be created automatically.<br/>If the flag is not specified, all backend config files will be written<br/>to the corresponding terraform component folders |       | no       |
| `--stacks`        | Only process the specified stacks (comma-separated values).<br/>The names of top-level stack config files and `atmos` stack names are supported                                                                                                                                                                                                                                                                                                                     |       | no       |
| `--components`    | Only generate backend files for the specified `atmos` components<br/>(comma-separated values)                                                                                                                                                                                                                                                                                                                                                                       |       | no       |
| `--format`        | Backend file format: `hcl`, `json`, `backend-config` (`hcl` is default)                                                                                                                                                                                                                                                                                                                                                                                             |       | no       |
| `--dry-run`       | Dry-run                                                                                                                                                                                                                                                                                                                                                                                                                                                             |       | no       |

<br/>

:::info
Refer to [Terraform backend configuration](https://developer.hashicorp.com/terraform/language/settings/backends/configuration) for more details
on `terraform` backends and supported formats
:::
