/**
 * Created on Wed Feb 22 18:55:13 2023
 * @author: Dmitry White
 */

/**
 * Create a function method that allow you to wrap an existing function.
 * The method signature would look something like this:
 * function speak(name) {
 *      return "Hello " + name;
 * }
 * speak = speak.wrap(function (original, yourName, myName) {
 *      greeting = original(yourName);
 *      return greeting + ", my name is " + myName;
 * });
 * var greeting = speak("Mary", "Kate");
 */

Function.prototype.wrap = function (callback) {
  return (...args) => callback(this, ...args);
};
