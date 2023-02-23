/**
 * Created on Wed Feb 22 18:21:49 2023
 * @author: Dmitry White
 */

/**
 * Your task is to create a change machine.
 * The machine accepts a range of specified coins and notes,
 * it returns change in 20p and 10p coins in the minimum amount of pieces.
 * A 50p for example would return two 20p pieces and one 10p piece.
 * The machine will always try and return change, if you input a 20p for example it will return "10p 10p".
 * The machine accepts these coins and notes: £5, £2, £1, 50p, 20p.
 * Any coins and notes which are not accepted by the machine will be returned.
 * If you were to put a £20 note into the machine for example,
 * it would be returned to you and not broken into change.
 * This change machine is programmed to accept and distribute strings rather than numbers.
 * The change will be returned as one string with the change separated by single spaces & no commas.
 * The values of the string will be descending.
 */

const MONEY_MAP = {
  FIVE_POUNDS: "£5",
  TWO_POUNDS: "£2",
  ONE_POUND: "£1",
  FIFTY_PENCE: "50p",
  TWENTY_PENCE: "20p",
  TEN_PENCE: "10p",
};

function changeMe(moneyIn) {
  const {
    FIVE_POUNDS,
    TWO_POUNDS,
    ONE_POUND,
    FIFTY_PENCE,
    TWENTY_PENCE,
    TEN_PENCE,
  } = MONEY_MAP;

  switch (moneyIn) {
    case FIVE_POUNDS:
      return `${TWENTY_PENCE} `.repeat(25).trim();
    case TWO_POUNDS:
      return `${TWENTY_PENCE} `.repeat(10).trim();
    case ONE_POUND:
      return `${TWENTY_PENCE} `.repeat(5).trim();
    case FIFTY_PENCE:
      return `${TWENTY_PENCE} ${TWENTY_PENCE} ${TEN_PENCE}`;
    case TWENTY_PENCE:
      return `${TEN_PENCE} ${TEN_PENCE}`;
    default:
      return moneyIn;
  }
}
