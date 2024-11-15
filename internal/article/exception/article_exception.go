package articleException

import (
	errorList "github.com/diki-haryadi/ztools/constant/error/error_list"
	customErrors "github.com/diki-haryadi/ztools/error/custom_error"
	errorUtils "github.com/diki-haryadi/ztools/error/error_utils"
)

func CreateArticleValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func ArticleBindingExc() error {
	articleBindingError := errorList.InternalErrorList.ArticleExceptions.BindingError
	return customErrors.NewBadRequestError(articleBindingError.Msg, articleBindingError.Code, nil)
}
