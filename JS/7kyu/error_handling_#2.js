/**
 * Created on Tue May 30 22:36:14 2023
 * @author: Dmitry White
 */

/**
 * TODO: You are provided to evaluate a string,
 * you must check for any HTML code (i.e. if any HTML tags are found),
 * if any code is found - return false,
 * if the input is not a string - throw a TypeError,
 * if the string is over 255 characters long or contains 0 characters - throw a RangeError
 * and if the string entered is null - throw a ReferenceError.
 */

const HTML_REGEX = /<(?:"[^"]*"['"]*|'[^']*'['"]*|[^'">])+>/g;

function validateMessage(msg) {
  switch (true) {
    case msg === null:
      throw new ReferenceError("Message is null!");
    case typeof msg !== "string" && !(msg instanceof String):
      throw new TypeError(
        `Message should be of type string but was of type ${typeof msg}!`
      );
    case msg.length === 0 || msg.length > 255:
      throw new RangeError(`Message contains ${msg.length} characters!`);
  }

  return !HTML_REGEX.test(msg);
}
