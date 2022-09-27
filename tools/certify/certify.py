#!/usr/bin/env python3
import subprocess
import sys
import tempfile

from validators import domain, ipv4, ipv6, slug


def remove_prefix(text, prefix):
    if text.startswith(prefix):
        return text[len(prefix) :]
    return text


def create_certs(certname, altnames, keysize=2048, days=3650):
    base_config = "basicConstraints=critical,CA:true,pathlen:0\nsubjectAltName="
    full_altnames = []
    for n in altnames:
        t = n
        if n.startswith("IP:"):
            t = remove_prefix(n, "IP:")
        if n.startswith("DNS:"):
            t = remove_prefix(n, "DNS:")
        good = domain(t) or ipv4(t) or ipv6(t) or slug(t)
        if bool(good) is False:
            raise OSError("Invalid domain passed")

        if "DNS:" not in n and "IP:" not in n:
            full_altnames.append(f"DNS:{n}")
        else:
            full_altnames.append(f"{n}")

    full_altnames = ",".join(full_altnames)
    base_config += full_altnames

    with tempfile.NamedTemporaryFile(mode="w+") as key, tempfile.NamedTemporaryFile(
        mode="w+"
    ) as csr, tempfile.NamedTemporaryFile(
        mode="w+"
    ) as cert, tempfile.NamedTemporaryFile(
        mode="w+"
    ) as ext:
        ext.write(base_config)
        ext.seek(0)

        subprocess.call(["openssl", "genrsa", "-out", f"{key.name}", f"{keysize}"])
        subprocess.call(
            [
                "openssl",
                "req",
                "-new",
                "-key",
                f"{key.name}",
                "-subj",
                f"/CN={certname}/",
                "-out",
                f"{csr.name}",
            ]
        )
        subprocess.call(
            [
                "openssl",
                "x509",
                "-req",
                "-signkey",
                f"{key.name}",
                "-days",
                f"{days}",
                "-in",
                f"{csr.name}",
                "-out",
                f"{cert.name}",
                "-extfile",
                f"{ext.name}",
            ]
        )
        subprocess.call(["openssl", "x509", "-in", f"{cert.name}", "-noout", "-text"])

        key.seek(0)
        cert.seek(0)
        return key.read(), cert.read()


if __name__ == "__main__":
    USAGE = """
certify.py
    Creates and self-signs X.509 SSL/TLS certificates with the "subjectAltName" extension.
    Allows for the specification of IP addresses as well as domain names.

Usage: ./certify.py example.com IP:127.0.0.1 DNS:localhost [www.example.com] [mail.example.com] [...]
"""
    try:
        certname = sys.argv[1]
        altnames = sys.argv[1:]
    except IndexError:
        print(USAGE)
        exit(1)

    key, cert = create_certs(certname, altnames)
    with open(f"{certname}.key", "w+") as keyfile, open(
        f"{certname}.crt", "w+"
    ) as certfile:
        keyfile.write(key)
        certfile.write(cert)
