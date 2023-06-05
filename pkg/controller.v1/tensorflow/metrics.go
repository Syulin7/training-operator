package tensorflow

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Define all the prometheus counters for all jobs
var (
	tfJobsCreatedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tf_operator_jobs_created",
			Help: "Counts number of TF jobs created",
		},
		[]string{"job_namespace", "job_name"},
	)
	tfJobsSuccessCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tf_operator_jobs_successful",
			Help: "Counts number of TF jobs successful",
		},
		[]string{"job_namespace", "job_name"},
	)
	tfJobsFailedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tf_operator_jobs_failed",
			Help: "Counts number of TF jobs failed",
		},
		[]string{"job_namespace", "job_name"},
	)
	tfJobsRestartCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tf_operator_jobs_restarted",
			Help: "Counts number of TF jobs restarted",
		},
		[]string{"job_namespace", "job_name"},
	)
	tfJobsDeletedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "tf_operator_jobs_deleted",
			Help: "Counts number of TF jobs deleted",
		},
		[]string{"job_namespace", "job_name"},
	)
	tfJobsWorkQueueLength = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tf_operator_work_queue_length",
			Help: "Length of TF jobs work queue",
		},
	)
)

func CreatedTFJobsCounterInc(namespace, name string) {
	tfJobsCreatedCount.WithLabelValues(namespace, name).Inc()
}

func SuccessfulTFJobsCounterInc(namespace, name string) {
	tfJobsSuccessCount.WithLabelValues(namespace, name).Inc()
}

func FailedTFJobsCounterInc(namespace, name string) {
	tfJobsFailedCount.WithLabelValues(namespace, name).Inc()
}

func RestartedTFJobsCounterInc(namespace, name string) {
	tfJobsRestartCount.WithLabelValues(namespace, name).Inc()
}

func DeletedTFJobsCounterInc(namespace, name string) {
	tfJobsDeletedCount.WithLabelValues(namespace, name).Inc()
}
