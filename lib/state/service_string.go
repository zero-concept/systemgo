// Code generated by "stringer -type=Service"; DO NOT EDIT

package state

import "fmt"

const _Service_name = "ServiceDeadServiceStartPreServiceStartServiceStartPostServiceRunningServiceExitedServiceReloadServiceStopServiceStopSigabrtServiceStopSigtermServiceStopSigkillServiceStopPostServiceFinalSigtermServiceFinalSigkillServiceFailedServiceAutoRestart"

var _Service_index = [...]uint8{0, 11, 26, 38, 54, 68, 81, 94, 105, 123, 141, 159, 174, 193, 212, 225, 243}

func (i Service) String() string {
	if i < 0 || i >= Service(len(_Service_index)-1) {
		return fmt.Sprintf("Service(%d)", i)
	}
	return _Service_name[_Service_index[i]:_Service_index[i+1]]
}
