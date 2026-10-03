package pokemon;

import move.PoisonPowder;
import move.SteelWing;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Type;

public class Charmander extends Pokemon {
  public Charmander(String name, int level) {
    super(name, level);
    this.setType(Type.FIRE);
    this.setStats(39, 52, 43, 60, 50, 65);
    this.setMove(
      new SteelWing(),
      new PoisonPowder()
    );
  }
}
