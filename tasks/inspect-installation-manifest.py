from cbor_diag import cbor2diag
import click

@click.command()
@click.argument('filename', type=click.File('rb'))
def run(filename):
    d = filename.read()
    c = int(d, 36)
    print(c)
    print(cbor2diag(c))

if __name__ == '__main__':
    run()
