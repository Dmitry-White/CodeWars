/**
 * Created on Wed Feb 22 18:11:37 2023
 * @author: Dmitry White
 */

/**
 * TODO: Complete the function that receives as input a string,
 * and produces outputs according to the menu.
 * A default case: if the input to the function is not any of the values in the table,
 * then the return value should be "Beer".
 * Make sure you cover the cases where certain words do not show up with correct capitalization.
 */

const MENU = {
  jabroni: "Patron Tequila",
  "school counselor": "Anything with Alcohol",
  programmer: "Hipster Craft Beer",
  "bike gang member": "Moonshine",
  politician: "Your tax dollars",
  rapper: "Cristal",
};

const DEFAULT_DRINK = "Beer";

function getDrinkByProfession(param) {
  const person = param.toLowerCase();

  return MENU[person] || DEFAULT_DRINK;
}
