package java

//
//import (
//	"errors"
//	corev1 "k8s.io/api/core/v1"
//	"reflect"
//	"strings"
//
//	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
//)
//
//// Config is the struct that contains the refined configuration for java instrumentation
//// it is used to pass the configuration to the mutator
//type Config struct {
//	Env              []corev1.EnvVar
//	Image            string
//	Endpoint         string
//	ServiceSource    string
//	Propagator       string
//	Metrics          string
//	Logs             string
//	TracerSampler    string
//	TracerSamplerArg string
//}
//
//// NewConfig get the instrumentation spec as parameter and generate Config from the object
//// 이 부분 굉장히 마음에 안 드는데 어떤 식으로 개선 할 지 고미좀 해보자...
//func NewConfig(spec v1.InstrumentationSpec) (*Config, error) {
//	res := Config{
//		Env: make([]corev1.EnvVar, 0, 10),
//	}
//	res.Image = spec.Java.Image
//
//	if len(spec.Java.EnvVars) != 0 {
//		for _, env := range spec.EnvVars {
//			res.Env = append(res.Env, corev1.EnvVar{
//				Name:  env.Name,
//				Value: env.Value,
//			})
//		}
//	} else if len(spec.EnvVars) != 0 {
//		for _, env := range spec.EnvVars {
//			res.Env = append(res.Env, corev1.EnvVar{
//				Name:  env.Name,
//				Value: env.Value,
//			})
//		}
//	}
//	if spec.Java.Endpoint != "" {
//		res.Endpoint = spec.Java.Endpoint
//	} else {
//		res.Endpoint = spec.Endpoint
//	}
//	if spec.Java.Sampling.SamplerArg != "" {
//		res.TracerSampler = spec.Java.Sampling.SamplerArg
//	} else {
//		res.TracerSampler = spec.Sampling.SamplerArg
//	}
//	if spec.Java.Sampling.Sampler != "" {
//		res.TracerSamplerArg = spec.Java.Sampling.Sampler
//	} else {
//		res.TracerSamplerArg = spec.Sampling.Sampler
//	}
//	if len(spec.Java.Propagator) != 0 {
//		res.Propagator = strings.Join(spec.Java.Propagator, ",")
//	} else {
//		res.Propagator = strings.Join(spec.Propagator, ",")
//	}
//	if spec.Java.Metrics != "" {
//		res.Metrics = spec.Java.Metrics
//	} else {
//		res.Metrics = spec.Metrics
//	}
//	if spec.Java.Logs != "" {
//		res.Logs = spec.Java.Logs
//	} else {
//		res.Logs = spec.Logs
//	}
//	if spec.Java.ServiceNameLabel != "" {
//		res.ServiceSource = spec.Java.ServiceNameLabel
//	}
//	return res.validate()
//}
//
//func (c *Config) validate() (*Config, error) {
//	// check if there is empty value in struct
//	s := reflect.ValueOf(c).Elem()
//	for i := 0; i < s.NumField(); i++ {
//		f := s.Field(i)
//		v := reflect.ValueOf(f.Interface())
//		if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
//			return nil, errors.New("invalid java configuration")
//		}
//	}
//	return c, nil
//}
