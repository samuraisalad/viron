viron-component-operation.ComponentOperation
  .ComponentOperation__head
    .ComponentOperation__title { opts.title }
    .ComponentOperation__description(if="{ !!opts.description }") { opts.description }
  .ComponentOperation__body
    viron-parameters(parameterObjects="{ opts.parameterObjects }" parameters="{ currentParameters }" additionalInfo="{ additionalInfo }" onChange="{ handleParametersChange }")
  .ComponentOperation__tail
    viron-button(label="{ submitButtonLabel }" type="{ submitButtonType }" onClick="{ handleSubmitButtonClick }")
    viron-button(label="閉じる" type="secondary" onClick="{ handleCancelButtonClick }")

  script.
    import '../../organisms/viron-parameters/index.tag';
    import '../../atoms/viron-button/index.tag';
    import script from './operation';
    this.external(script);