function jsid64(length) {
    length = length || 21;
    var crypto = self.crypto || self.msCrypto || {
        getRandomValues: array => {
            for (let i = 0, l = array.length; i < l; i++) {
                array[i] = Math.floor(Math.random() * 256);
            }
            return array;
        }
    };
    var alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~';
    var bytes = new Int8Array(length);
    crypto.getRandomValues(bytes);
    var result = "";
    for(var i = 0; i < bytes.length; i++) {
        result = result + alphabet[bytes[i]&63];
    }

    return result;
}
