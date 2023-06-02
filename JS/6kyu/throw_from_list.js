/**
 * Created on Fri Jun 02 19:31:14 2023
 * @author: Dmitry White
 */

/**
 * TODO: Your task is to throw errors from the preloaded list of errors.
 * All your solution should do is to throw errors if the given input is invalid.
 * The function you have to write is for username and password validation
 * which must adhere to certain requirements and throw relevant errors when necessary.
 * The order of validation must follow the order of errors in the `Errors To Throw` section.
 *
 * Criteria
 * The `username` must follow these rules:
 * - Maxiumum 12 characters long
 * - No special characters like `(){}[]|;:'"/?.,<>~-=+*&^%$@!`
 * - Spaces at the start and end of names must be trimmed
 * - Must be at least 1 character long
 *
 * The `password` must follow these rules:
 * - Maxiumum 24 characters long
 * - Must be at least 8 character long
 * - These allowed special characters `;:?.,<>~*^%$ @!_`
 * - Spaces at the start and end of passwords must be trimmed
 * - Must contain at least 1 capital letter and 1 numeric number
 * - Password must not contain the username but only if case matches
 *
 * Errors to throw
 * There is a `const` that is preloaded into the solution called `ERRORS`.
 * This is an object that contain error properties and proper error logs
 * which tell us what exactly went wrong.
 * You must implement this in your solution throwing these errors only:
 * - `usernameTooLong` This should get thrown when the username is over 12 characters long
 * - `usernameTooShort` This should get thrown when the username is less than 1 character long
 * - `usernameInvalidCharacters` This should get thrown when the username contains illegal characters
 * - `passwordTooLong` This should get thrown when the password is over 24 characters long
 * - `passwordTooShort` This should get thrown when the password is less than 8 characters
 * - `passwordInvalidCharacters` This should get thrown whent he password contains any illegal characters
 * - `passwordNoCapital` This should get thrown when the password contains no Capital letters
 * - `passwordNoNumber` This should get thrown when the password does not contain any numeric values
 * - `passwordContainsUsername` This should get thrown when the username is present in the password but only if case matches
 * Usage:
 * ```
 *  ERRORS.usernameTooLong(username);
 *  ERRORS.passwordTooLong(password);
 * ```
 */
const isUsernameTooLong = (username) => username.length > 12;
const isUsernameTooShort = (username) => username.length < 1;
const isUsernameInvalidCharacters = (username) => {
  const invalidCharacters = /[(){}[\]|;:'"\/?.,<>~\-=+\*&^%$@!]/;
  return invalidCharacters.test(username);
};

const isPasswordTooLong = (password) => password.length > 24;
const isPasswordTooShort = (password) => password.length < 8;
const isPasswordInvalidCharacters = (password) => {
  const validCharacters = /^[a-zA-Z0-9_;:?.,<>~*^%$ @!_]*$/;
  return !validCharacters.test(password);
};
const isPasswordNoCapital = (password) => {
  const capitalRegex = /[A-Z]/;
  return !capitalRegex.test(password);
};
const isPasswordNoNumber = (password) => {
  const numberRegex = /[0-9]/;
  return !numberRegex.test(password);
};

const isPasswordContainsUsername = (username, password) => {
  const usernameRegex = new RegExp(`^.*${username}.*$`);
  return usernameRegex.test(password);
};

const CHECK_MAP = {
  isUsernameTooLong: ERRORS.usernameTooLong,
  isUsernameTooShort: ERRORS.usernameTooShort,
  isUsernameInvalidCharacters: ERRORS.usernameInvalidCharacters,
  isPasswordTooLong: ERRORS.passwordTooLong,
  isPasswordTooShort: ERRORS.passwordTooShort,
  isPasswordInvalidCharacters: ERRORS.passwordInvalidCharacters,
  isPasswordNoCapital: ERRORS.passwordNoCapital,
  isPasswordNoNumber: ERRORS.passwordNoNumber,
  isPasswordContainsUsername: ERRORS.passwordContainsUsername,
};

const usernameChecks = [
  isUsernameTooLong,
  isUsernameTooShort,
  isUsernameInvalidCharacters,
];

const passwordChecks = [
  isPasswordTooLong,
  isPasswordTooShort,
  isPasswordInvalidCharacters,
  isPasswordNoCapital,
  isPasswordNoNumber,
];

const relatedChecks = [isPasswordContainsUsername];

const checkTerm = (checks, term) => {
  const failingCheck = checks.find((violate) => violate(term));
  if (failingCheck) {
    throw CHECK_MAP[failingCheck.name](term);
  }
  return true;
};

const checkRelated = (checks, username, password) => {
  const failingCheck = checks.find((violate) => violate(username, password));
  if (failingCheck) {
    throw CHECK_MAP[failingCheck.name](password);
  }
  return true;
};

function validate(username, password) {
  const checkResults = [
    checkTerm(usernameChecks, username.trim()),
    checkTerm(passwordChecks, password.trim()),
    checkRelated(relatedChecks, username.trim(), password.trim()),
  ];
  return checkResults.every((check) => check);
}
