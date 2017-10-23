// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v2alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	v1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v2alpha1_CronJob_To_batch_CronJob,
		Convert_batch_CronJob_To_v2alpha1_CronJob,
		Convert_v2alpha1_CronJobList_To_batch_CronJobList,
		Convert_batch_CronJobList_To_v2alpha1_CronJobList,
		Convert_v2alpha1_CronJobSpec_To_batch_CronJobSpec,
		Convert_batch_CronJobSpec_To_v2alpha1_CronJobSpec,
		Convert_v2alpha1_CronJobStatus_To_batch_CronJobStatus,
		Convert_batch_CronJobStatus_To_v2alpha1_CronJobStatus,
		Convert_v2alpha1_Job_To_batch_Job,
		Convert_batch_Job_To_v2alpha1_Job,
		Convert_v2alpha1_JobCondition_To_batch_JobCondition,
		Convert_batch_JobCondition_To_v2alpha1_JobCondition,
		Convert_v2alpha1_JobList_To_batch_JobList,
		Convert_batch_JobList_To_v2alpha1_JobList,
		Convert_v2alpha1_JobSpec_To_batch_JobSpec,
		Convert_batch_JobSpec_To_v2alpha1_JobSpec,
		Convert_v2alpha1_JobStatus_To_batch_JobStatus,
		Convert_batch_JobStatus_To_v2alpha1_JobStatus,
		Convert_v2alpha1_JobTemplate_To_batch_JobTemplate,
		Convert_batch_JobTemplate_To_v2alpha1_JobTemplate,
		Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec,
		Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec,
	)
}

func autoConvert_v2alpha1_CronJob_To_batch_CronJob(in *CronJob, out *batch.CronJob, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_CronJobSpec_To_batch_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_CronJobStatus_To_batch_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_CronJob_To_batch_CronJob(in *CronJob, out *batch.CronJob, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJob_To_batch_CronJob(in, out, s)
}

func autoConvert_batch_CronJob_To_v2alpha1_CronJob(in *batch.CronJob, out *CronJob, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_CronJobSpec_To_v2alpha1_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_CronJobStatus_To_v2alpha1_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_CronJob_To_v2alpha1_CronJob(in *batch.CronJob, out *CronJob, s conversion.Scope) error {
	return autoConvert_batch_CronJob_To_v2alpha1_CronJob(in, out, s)
}

func autoConvert_v2alpha1_CronJobList_To_batch_CronJobList(in *CronJobList, out *batch.CronJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.CronJob, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_CronJob_To_batch_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v2alpha1_CronJobList_To_batch_CronJobList(in *CronJobList, out *batch.CronJobList, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobList_To_batch_CronJobList(in, out, s)
}

func autoConvert_batch_CronJobList_To_v2alpha1_CronJobList(in *batch.CronJobList, out *CronJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronJob, len(*in))
		for i := range *in {
			if err := Convert_batch_CronJob_To_v2alpha1_CronJob(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_batch_CronJobList_To_v2alpha1_CronJobList(in *batch.CronJobList, out *CronJobList, s conversion.Scope) error {
	return autoConvert_batch_CronJobList_To_v2alpha1_CronJobList(in, out, s)
}

func autoConvert_v2alpha1_CronJobSpec_To_batch_CronJobSpec(in *CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = batch.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_CronJobSpec_To_batch_CronJobSpec(in *CronJobSpec, out *batch.CronJobSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobSpec_To_batch_CronJobSpec(in, out, s)
}

func autoConvert_batch_CronJobSpec_To_v2alpha1_CronJobSpec(in *batch.CronJobSpec, out *CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_CronJobSpec_To_v2alpha1_CronJobSpec(in *batch.CronJobSpec, out *CronJobSpec, s conversion.Scope) error {
	return autoConvert_batch_CronJobSpec_To_v2alpha1_CronJobSpec(in, out, s)
}

func autoConvert_v2alpha1_CronJobStatus_To_batch_CronJobStatus(in *CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]api.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*v1.Time)(unsafe.Pointer(in.LastScheduleTime))
	return nil
}

func Convert_v2alpha1_CronJobStatus_To_batch_CronJobStatus(in *CronJobStatus, out *batch.CronJobStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobStatus_To_batch_CronJobStatus(in, out, s)
}

func autoConvert_batch_CronJobStatus_To_v2alpha1_CronJobStatus(in *batch.CronJobStatus, out *CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]api_v1.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*v1.Time)(unsafe.Pointer(in.LastScheduleTime))
	return nil
}

func Convert_batch_CronJobStatus_To_v2alpha1_CronJobStatus(in *batch.CronJobStatus, out *CronJobStatus, s conversion.Scope) error {
	return autoConvert_batch_CronJobStatus_To_v2alpha1_CronJobStatus(in, out, s)
}

func autoConvert_v2alpha1_Job_To_batch_Job(in *Job, out *batch.Job, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobStatus_To_batch_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_Job_To_batch_Job(in *Job, out *batch.Job, s conversion.Scope) error {
	return autoConvert_v2alpha1_Job_To_batch_Job(in, out, s)
}

func autoConvert_batch_Job_To_v2alpha1_Job(in *batch.Job, out *Job, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobSpec_To_v2alpha1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_batch_JobStatus_To_v2alpha1_JobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_Job_To_v2alpha1_Job(in *batch.Job, out *Job, s conversion.Scope) error {
	return autoConvert_batch_Job_To_v2alpha1_Job(in, out, s)
}

func autoConvert_v2alpha1_JobCondition_To_batch_JobCondition(in *JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	out.Type = batch.JobConditionType(in.Type)
	out.Status = api.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_v2alpha1_JobCondition_To_batch_JobCondition(in *JobCondition, out *batch.JobCondition, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobCondition_To_batch_JobCondition(in, out, s)
}

func autoConvert_batch_JobCondition_To_v2alpha1_JobCondition(in *batch.JobCondition, out *JobCondition, s conversion.Scope) error {
	out.Type = JobConditionType(in.Type)
	out.Status = api_v1.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func Convert_batch_JobCondition_To_v2alpha1_JobCondition(in *batch.JobCondition, out *JobCondition, s conversion.Scope) error {
	return autoConvert_batch_JobCondition_To_v2alpha1_JobCondition(in, out, s)
}

func autoConvert_v2alpha1_JobList_To_batch_JobList(in *JobList, out *batch.JobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]batch.Job, len(*in))
		for i := range *in {
			if err := Convert_v2alpha1_Job_To_batch_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v2alpha1_JobList_To_batch_JobList(in *JobList, out *batch.JobList, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobList_To_batch_JobList(in, out, s)
}

func autoConvert_batch_JobList_To_v2alpha1_JobList(in *batch.JobList, out *JobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Job, len(*in))
		for i := range *in {
			if err := Convert_batch_Job_To_v2alpha1_Job(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_batch_JobList_To_v2alpha1_JobList(in *batch.JobList, out *JobList, s conversion.Scope) error {
	return autoConvert_batch_JobList_To_v2alpha1_JobList(in, out, s)
}

func autoConvert_v2alpha1_JobSpec_To_batch_JobSpec(in *JobSpec, out *batch.JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := api_v1.Convert_v1_PodTemplateSpec_To_api_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_batch_JobSpec_To_v2alpha1_JobSpec(in *batch.JobSpec, out *JobSpec, s conversion.Scope) error {
	out.Parallelism = (*int32)(unsafe.Pointer(in.Parallelism))
	out.Completions = (*int32)(unsafe.Pointer(in.Completions))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.ManualSelector = (*bool)(unsafe.Pointer(in.ManualSelector))
	if err := api_v1.Convert_api_PodTemplateSpec_To_v1_PodTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v2alpha1_JobStatus_To_batch_JobStatus(in *JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]batch.JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*v1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*v1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func Convert_v2alpha1_JobStatus_To_batch_JobStatus(in *JobStatus, out *batch.JobStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobStatus_To_batch_JobStatus(in, out, s)
}

func autoConvert_batch_JobStatus_To_v2alpha1_JobStatus(in *batch.JobStatus, out *JobStatus, s conversion.Scope) error {
	out.Conditions = *(*[]JobCondition)(unsafe.Pointer(&in.Conditions))
	out.StartTime = (*v1.Time)(unsafe.Pointer(in.StartTime))
	out.CompletionTime = (*v1.Time)(unsafe.Pointer(in.CompletionTime))
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func Convert_batch_JobStatus_To_v2alpha1_JobStatus(in *batch.JobStatus, out *JobStatus, s conversion.Scope) error {
	return autoConvert_batch_JobStatus_To_v2alpha1_JobStatus(in, out, s)
}

func autoConvert_v2alpha1_JobTemplate_To_batch_JobTemplate(in *JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_JobTemplate_To_batch_JobTemplate(in *JobTemplate, out *batch.JobTemplate, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplate_To_batch_JobTemplate(in, out, s)
}

func autoConvert_batch_JobTemplate_To_v2alpha1_JobTemplate(in *batch.JobTemplate, out *JobTemplate, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_JobTemplate_To_v2alpha1_JobTemplate(in *batch.JobTemplate, out *JobTemplate, s conversion.Scope) error {
	return autoConvert_batch_JobTemplate_To_v2alpha1_JobTemplate(in, out, s)
}

func autoConvert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in *JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v2alpha1_JobSpec_To_batch_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in *JobTemplateSpec, out *batch.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplateSpec_To_batch_JobTemplateSpec(in, out, s)
}

func autoConvert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *batch.JobTemplateSpec, out *JobTemplateSpec, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_batch_JobSpec_To_v2alpha1_JobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func Convert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *batch.JobTemplateSpec, out *JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_batch_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in, out, s)
}
