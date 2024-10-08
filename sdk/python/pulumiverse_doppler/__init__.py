# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .branch_config import *
from .environment import *
from .get_secrets import *
from .get_user import *
from .group import *
from .group_member import *
from .group_members import *
from .project import *
from .project_role import *
from .provider import *
from .secret import *
from .service_account import *
from .service_account_token import *
from .service_token import *
from .webhook import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_doppler.config as __config
    config = __config
    import pulumiverse_doppler.integration as __integration
    integration = __integration
    import pulumiverse_doppler.projectmember as __projectmember
    projectmember = __projectmember
    import pulumiverse_doppler.secretssync as __secretssync
    secretssync = __secretssync
else:
    config = _utilities.lazy_import('pulumiverse_doppler.config')
    integration = _utilities.lazy_import('pulumiverse_doppler.integration')
    projectmember = _utilities.lazy_import('pulumiverse_doppler.projectmember')
    secretssync = _utilities.lazy_import('pulumiverse_doppler.secretssync')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "doppler",
  "mod": "index/branchConfig",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/branchConfig:BranchConfig": "BranchConfig"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/environment",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/environment:Environment": "Environment"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/group",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/group:Group": "Group"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/groupMember",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/groupMember:GroupMember": "GroupMember"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/groupMembers",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/groupMembers:GroupMembers": "GroupMembers"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/project",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/project:Project": "Project"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/projectRole",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/projectRole:ProjectRole": "ProjectRole"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/secret",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/secret:Secret": "Secret"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/serviceAccount",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/serviceAccount:ServiceAccount": "ServiceAccount"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/serviceAccountToken",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/serviceAccountToken:ServiceAccountToken": "ServiceAccountToken"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/serviceToken",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/serviceToken:ServiceToken": "ServiceToken"
  }
 },
 {
  "pkg": "doppler",
  "mod": "index/webhook",
  "fqn": "pulumiverse_doppler",
  "classes": {
   "doppler:index/webhook:Webhook": "Webhook"
  }
 },
 {
  "pkg": "doppler",
  "mod": "integration/awsParameterStore",
  "fqn": "pulumiverse_doppler.integration",
  "classes": {
   "doppler:integration/awsParameterStore:AwsParameterStore": "AwsParameterStore"
  }
 },
 {
  "pkg": "doppler",
  "mod": "integration/awsSecretsManager",
  "fqn": "pulumiverse_doppler.integration",
  "classes": {
   "doppler:integration/awsSecretsManager:AwsSecretsManager": "AwsSecretsManager"
  }
 },
 {
  "pkg": "doppler",
  "mod": "integration/circleci",
  "fqn": "pulumiverse_doppler.integration",
  "classes": {
   "doppler:integration/circleci:Circleci": "Circleci"
  }
 },
 {
  "pkg": "doppler",
  "mod": "integration/flyio",
  "fqn": "pulumiverse_doppler.integration",
  "classes": {
   "doppler:integration/flyio:Flyio": "Flyio"
  }
 },
 {
  "pkg": "doppler",
  "mod": "integration/terraformCloud",
  "fqn": "pulumiverse_doppler.integration",
  "classes": {
   "doppler:integration/terraformCloud:TerraformCloud": "TerraformCloud"
  }
 },
 {
  "pkg": "doppler",
  "mod": "projectMember/group",
  "fqn": "pulumiverse_doppler.projectmember",
  "classes": {
   "doppler:projectMember/group:Group": "Group"
  }
 },
 {
  "pkg": "doppler",
  "mod": "projectMember/serviceAccount",
  "fqn": "pulumiverse_doppler.projectmember",
  "classes": {
   "doppler:projectMember/serviceAccount:ServiceAccount": "ServiceAccount"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/awsParameterStore",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/awsParameterStore:AwsParameterStore": "AwsParameterStore"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/awsSecretsManager",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/awsSecretsManager:AwsSecretsManager": "AwsSecretsManager"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/circleci",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/circleci:Circleci": "Circleci"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/flyio",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/flyio:Flyio": "Flyio"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/githubActions",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/githubActions:GithubActions": "GithubActions"
  }
 },
 {
  "pkg": "doppler",
  "mod": "secretsSync/terraformCloud",
  "fqn": "pulumiverse_doppler.secretssync",
  "classes": {
   "doppler:secretsSync/terraformCloud:TerraformCloud": "TerraformCloud"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "doppler",
  "token": "pulumi:providers:doppler",
  "fqn": "pulumiverse_doppler",
  "class": "Provider"
 }
]
"""
)
