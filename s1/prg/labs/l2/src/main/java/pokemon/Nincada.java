package pokemon;

import move.DoubleTeam;
import move.FakeOut;
import move.MuddyWater;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Type;

final public class Nincada extends Pokemon {
  public Nincada(String name, int level) {
    super(name, level);
    this.setType(Type.BUG, Type.GROUND);
    this.setStats(31, 45, 90, 30, 30, 40);
    this.setMove(
      new MuddyWater(),
      new FakeOut(),
      new DoubleTeam()
    );
  }
}
