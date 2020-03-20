/*
 * NSSF Configuration Factory
 */

package factory

import (
	. "gofree5gc/lib/openapi/models"
)

type Subscription struct {
	SubscriptionId string `yaml:"subscriptionId"`

	SubscriptionData *NssfEventSubscriptionCreateData `yaml:"subscriptionData"`
}
