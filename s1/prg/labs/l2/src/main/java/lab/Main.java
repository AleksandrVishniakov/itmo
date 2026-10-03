package lab;

import ru.ifmo.se.pokemon.Battle;
import ru.ifmo.se.pokemon.Pokemon;
import pokemon.Charizard;
import pokemon.Charmander;
import pokemon.Charmeleon;
import pokemon.Gorebyss;
import pokemon.Nincada;
import pokemon.Scizor;

public class Main {
  static public void main(String[] args) {
    Battle b = new Battle();

    Pokemon gorebyss = new Gorebyss("Gorebyss", 1);
    Pokemon nincada = new Nincada("Nincada", 1);
    Pokemon scizor = new Scizor("Scizor", 1);

    Pokemon charmander = new Charmander("Charmander", 1);
    Pokemon charmeleon = new Charmeleon("Charmeleon", 1);
    Pokemon charizard = new Charizard("Charizard", 1);

    b.addAlly(charmander);
    b.addAlly(charmeleon);
    b.addAlly(charizard);

    b.addFoe(gorebyss);
    b.addFoe(nincada);
    b.addFoe(scizor);

    b.go();
  }
}
