// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CronJobSpec CronJobSpec describes how the job execution will look like and when it will actually run.
//
// swagger:model CronJobSpec
type CronJobSpec struct {

	// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	//
	//
	ConcurrencyPolicy string `json:"concurrencyPolicy,omitempty"`

	// The number of failed finished jobs to retain. Value must be non-negative integer. Defaults to 1.
	FailedJobsHistoryLimit int32 `json:"failedJobsHistoryLimit,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	// Required: true
	JobTemplate *JobTemplateSpec `json:"jobTemplate"`

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	// Required: true
	Schedule *string `json:"schedule"`

	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds int64 `json:"startingDeadlineSeconds,omitempty"`

	// The number of successful finished jobs to retain. Value must be non-negative integer. Defaults to 3.
	SuccessfulJobsHistoryLimit int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend bool `json:"suspend,omitempty"`

	// The time zone for the given schedule, see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. If not specified, this will rely on the time zone of the kube-controller-manager process. ALPHA: This field is in alpha and must be enabled via the `CronJobTimeZone` feature gate.
	TimeZone string `json:"timeZone,omitempty"`
}
