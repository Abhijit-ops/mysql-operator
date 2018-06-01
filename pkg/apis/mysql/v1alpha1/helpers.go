// Copyright 2018 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/oracle/mysql-operator/pkg/constants"
	"github.com/oracle/mysql-operator/pkg/version"
)

// EnsureDefaults will ensure that if a user omits and fields in the
// spec that are required, we set some sensible defaults.
// For example a user can choose to omit the version
// and number of replics
func (c *Cluster) EnsureDefaults() *Cluster {
	if c.Spec.Replicas == 0 {
		c.Spec.Replicas = defaultReplicas
	}

	if c.Spec.BaseServerID == 0 {
		c.Spec.BaseServerID = defaultBaseServerID
	}

	if c.Spec.Version == "" {
		c.Spec.Version = defaultVersion
	}

	return c
}

// Validate returns an error if a cluster is invalid
func (c *Cluster) Validate() error {
	return validateCluster(c).ToAggregate()
}

// RequiresConfigMount will return true if a user has specified a config map
// for configuring the cluster else false
func (c *Cluster) RequiresConfigMount() bool {
	return c.Spec.ConfigRef != nil
}

// RequiresSecret returns true if a secret should be generated
// for a MySQL cluster else false
func (c *Cluster) RequiresSecret() bool {
	return c.Spec.SecretRef == nil
}

// RequiresCustomSSLSetup returns true is the user has provided a secret
// that contains CA cert, server cert and server key for group replication
// SSL support
func (c *Cluster) RequiresCustomSSLSetup() bool {
	return c.Spec.SSLSecretRef != nil
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (b Backup) EnsureDefaults() *Backup {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if b.Labels == nil {
			b.Labels = make(map[string]string)
		}
		_, hasKey := b.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			SetOperatorVersionLabel(b.Labels, buildVersion)
		}
	}
	return &b
}

// Validate checks if the resource spec is valid.
func (b Backup) Validate() error {
	return validateBackup(&b).ToAggregate()
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (b BackupSchedule) EnsureDefaults() *BackupSchedule {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if b.Labels == nil {
			b.Labels = make(map[string]string)
		}
		_, hasKey := b.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			SetOperatorVersionLabel(b.Labels, buildVersion)
		}
	}
	return &b
}

// Validate checks if the resource spec is valid.
func (b BackupSchedule) Validate() error {
	return validateBackupSchedule(&b).ToAggregate()
}

// EnsureDefaults can be invoked to ensure the default values are present.
func (r Restore) EnsureDefaults() *Restore {
	buildVersion := version.GetBuildVersion()
	if buildVersion != "" {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		_, hasKey := r.Labels[constants.MySQLOperatorVersionLabel]
		if !hasKey {
			SetOperatorVersionLabel(r.Labels, buildVersion)
		}
	}
	return &r
}

// Validate checks if the resource spec is valid.
func (r Restore) Validate() error {
	return validateRestore(&r).ToAggregate()
}
