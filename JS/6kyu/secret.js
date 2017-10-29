/**
 * Created on Sun Oct 29 21:18:52 2017
 * @author: Dmitry White
 */

 /**
  * TODO: Implement a function createSecretHolder(secret)
  * which accepts any value as secret and returns an object
  * with ONLY two methods:
  *     getSecret() which returns the secret
  *     setSecret() which sets the secret
  */

function createSecretHolder(secret) {
    return {
        getSecret: function() { return secret; },
        setSecret: function(new_secret) { secret = new_secret; }
    };
};
