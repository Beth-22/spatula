// plugins/vee-validate.js
import { defineRule, configure } from "vee-validate";
import * as rules from "@vee-validate/rules";

// Define all rules
Object.keys(rules).forEach((rule) => {
  defineRule(rule, rules[rule]);
});

// Configure VeeValidate
configure({
  generateMessage: (context) => {
    const messages = {
      required: "This field is required",
      email: "This field must be a valid email",
      min: `This field must be at least ${context.rule.params} characters`,
      max: `This field must be less than ${context.rule.params} characters`,
    };
    return (
      messages[context.rule.name] || `The field ${context.field} is invalid`
    );
  },
});
