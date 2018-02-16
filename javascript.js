function jsid64(length) {
    length = length || 21;
    var crypto = self.crypto || self.msCrypto;
    var alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~';
    var bytes = new Int8Array(length);
    crypto.getRandomValues(bytes);
    var result = "";
    for(var i = 0; i < bytes.length; i++) {
        result = result + alphabet[bytes[i]&63];
    }

    return result;
}
