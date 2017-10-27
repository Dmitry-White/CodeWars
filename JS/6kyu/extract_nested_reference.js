/**
 * Created on Fri Oct 27 22:06:47 2017
 * @author: Dmitry White
 */

 /**
  * TODO: You are given a complex object that has many deeply nested variables.
  * You don't want to go the usual if obj.property == null route.
  * Create a prototype method that given a nested path, either return
  * the value or undefined.
  */

Object.prototype.hash = function(string) {
    let obj = this;
    let fields = string.split(".");
    let len = fields.length;
    for (let i = 0; i<len; i++){
        if (fields[i] in obj) {
          obj = obj[fields[i]];
        }else return undefined;
    }
    return obj;
};
