package resource

import (
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	ssmSvc := ssm.New(req.Session)

	res, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(*currentModel.ParameterName.Value()),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	currentModel.Value = encoding.NewString(*res.Parameter.Value)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	ssmSvc := ssm.New(req.Session)

	res, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(*currentModel.ParameterName.Value()),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	currentModel.Value = encoding.NewString(*res.Parameter.Value)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	ssmSvc := ssm.New(req.Session)

	res, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(*currentModel.ParameterName.Value()),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		return handler.ProgressEvent{}, err
	}

	currentModel.Value = encoding.NewString(*res.Parameter.Value)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// no-op

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "no-op",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// no-op

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "no-op",
		ResourceModel:   currentModel,
	}, nil
}
