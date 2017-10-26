/**
 * Created on Thu Oct 26 21:30:28 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Define a class Person with the following properties:
  *     A constructor that accepts 4 arguments: firstName/FirstName
  *      (defaults to "John" if not set), lastName/LastName
  *      (defaults to "Doe" if not set), age/Age
  *      (defaults to 0 if not set) and gender/Gender
  *      (defaults to "Male" if not set).
  *     These should be stored in this.firstName/this.FirstName,
  *     this.lastName/this.LastName, this.age/this.Age and
  *     this.gender/this.Gender respectively.
  *     A method sayFullName/SayFullName that accepts no arguments and
  *      returns the full name (e.g. "John Doe")
  *     A class/static method greetExtraTerrestrials/GreetExtraTerrestrials
  *      that accepts one parameter raceName and
  *      returns "Welcome to Planet Earth raceName".
  *      For example, if the race name is "Martians",
  *      it should say "Welcome to Planet Earth Martians"
  */


class Person {
    constructor (firstName, lastName, age, gender) {
        firstName !== undefined ? this.firstName = firstName : this.firstName = "John";
        lastName !== undefined ? this.lastName = lastName : this.lastName = "Doe";
        age !== undefined ? this.age = age : this.age = 0;
        gender !== undefined ? this.gender = gender : this.gender = "Male";
    }
    sayFullName () {
        return this.firstName + " " + this.lastName;
    }
    static greetExtraTerrestrials (raceName) {
        return "Welcome to Planet Earth " + raceName;
    }
  }
